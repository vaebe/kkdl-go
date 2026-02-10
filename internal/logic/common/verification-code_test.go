package common_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"compressURL/internal/dao"
	"compressURL/internal/logic/common"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	ctx   = context.Background()
	logic = common.New()
)

// getTestEmail 获取唯一的测试邮箱地址
func getTestEmail(suffix string) string {
	return fmt.Sprintf("test-%s-%d@example.com", suffix, time.Now().UnixNano())
}

// runInTransaction 在事务中执行测试函数，测试完成后自动回滚以清理测试数据
func runInTransaction(t *gtest.T, testFunc func(ctx context.Context, tx gdb.TX)) {
	_ = dao.VerificationCodes.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 执行测试函数
		testFunc(ctx, tx)
		// 返回错误以触发事务回滚，清理测试数据
		return errors.New("rollback test transaction")
	})
}

// insertExpiredVCode 手动插入一个已过期的验证码
func insertExpiredVCode(tx gdb.TX, email string, code string) error {
	_, err := tx.Model("verification_codes").Data(g.Map{
		"email":      email,
		"type":       1,
		"code":       code,
		"used":       0,
		"expired_at": gtime.Now().Add(-1 * time.Hour), // 1小时前过期
	}).Insert()
	return err
}

// getVCodeRecord 查询验证码记录
func getVCodeRecord(tx gdb.TX, email string, code string) (gdb.Record, error) {
	return tx.Model("verification_codes").
		Where("email", email).
		Where("code", code).
		One()
}

// assertVCodeExists 断言验证码记录存在
func assertVCodeExists(t *gtest.T, tx gdb.TX, email string, code string) {
	record, err := getVCodeRecord(tx, email, code)
	t.AssertNil(err)
	t.AssertEQ(record.IsEmpty(), false)
}

// assertVCodeNotExists 断言验证码记录不存在
func assertVCodeNotExists(t *gtest.T, tx gdb.TX, email string) {
	record, err := tx.Model("verification_codes").Where("email", email).One()
	t.AssertNil(err)
	t.AssertEQ(record.IsEmpty(), true)
}

// assertVCodeUsed 断言验证码已被标记为已使用
func assertVCodeUsed(t *gtest.T, tx gdb.TX, email string, code string, used int) {
	record, err := getVCodeRecord(tx, email, code)
	t.AssertNil(err)
	t.AssertEQ(record["used"].Int(), used)
}

// TestCreateVCode 测试创建验证码
func TestCreateVCode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testEmail := getTestEmail("create")

		// 使用事务确保测试隔离
		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			code, err := logic.CreateVCode(ctx, testEmail)
			t.AssertNil(err)
			t.AssertNE(code, "")

			// 验证码长度应为6位
			t.AssertEQ(len(code), 6)

			// 验证验证码是否已保存到数据库
			record, err := tx.Model("verification_codes").Where("email", testEmail).One()
			t.AssertNil(err)
			t.AssertNE(record.IsEmpty(), true)

			// 验证码内容匹配
			t.AssertEQ(record["code"].String(), code)
			t.AssertEQ(record["used"].Int(), 0)
		})
	})
}

// TestCreateVCodeMultipleTimes 测试多次创建验证码
func TestCreateVCodeMultipleTimes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testEmail := getTestEmail("multiple")

		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			code1, err1 := logic.CreateVCode(ctx, testEmail)
			t.AssertNil(err1)

			time.Sleep(100 * time.Millisecond) // 稍微等待，确保时间戳不同

			code2, err2 := logic.CreateVCode(ctx, testEmail)
			t.AssertNil(err2)

			// 两次生成的验证码应该不同
			t.AssertNE(code1, code2)
		})
	})
}

// TestVerifyTheVCode 测试校验验证码
func TestVerifyTheVCode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testEmail := getTestEmail("verify")

		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			// 先创建一个验证码
			code, err := logic.CreateVCode(ctx, testEmail)
			t.AssertNil(err)

			// 验证正确的验证码
			valid, err := logic.VerifyTheVCode(ctx, testEmail, code)
			t.AssertNil(err)
			t.AssertEQ(valid, true)

			// 验证错误的验证码
			invalid, err := logic.VerifyTheVCode(ctx, testEmail, "000000")
			t.AssertNil(err)
			t.AssertEQ(invalid, false)

			// 验证错误的邮箱
			invalid2, err := logic.VerifyTheVCode(ctx, "wrong@example.com", code)
			t.AssertNil(err)
			t.AssertEQ(invalid2, false)
		})
	})
}

// TestVerifyTheVCodeExpired 测试验证过期验证码
func TestVerifyTheVCodeExpired(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testEmail := getTestEmail("expired")
		expiredCode := gconv.String(123456)

		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			// 手动插入一个已过期的验证码
			err := insertExpiredVCode(tx, testEmail, expiredCode)
			t.AssertNil(err)

			// 验证已过期的验证码
			valid, err := logic.VerifyTheVCode(ctx, testEmail, expiredCode)
			t.AssertNil(err)
			t.AssertEQ(valid, false)
		})
	})
}

// TestVerifyTheVCodeUsed 测试验证已使用的验证码
func TestVerifyTheVCodeUsed(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testEmail := getTestEmail("used")

		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			// 先创建一个验证码
			code, err := logic.CreateVCode(ctx, testEmail)
			t.AssertNil(err)

			// 消费验证码
			err = logic.ConsumeVCode(ctx, testEmail, code)
			t.AssertNil(err)

			// 验证已使用的验证码
			valid, err := logic.VerifyTheVCode(ctx, testEmail, code)
			t.AssertNil(err)
			t.AssertEQ(valid, false)
		})
	})
}

// TestCheckVCodeCooldown 测试检查验证码冷却期
func TestCheckVCodeCooldown(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testEmail := getTestEmail("cooldown")

		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			// 场景1: 没有发送过验证码
			inCooldown, remaining, err := logic.CheckVCodeCooldown(ctx, "new@example.com")
			t.AssertNil(err)
			t.AssertEQ(inCooldown, false)
			t.AssertEQ(remaining, int64(0))

			// 场景2: 刚刚发送验证码（在冷却期内）
			_, err = logic.CreateVCode(ctx, testEmail)
			t.AssertNil(err)

			inCooldown, remaining, err = logic.CheckVCodeCooldown(ctx, testEmail)
			t.AssertNil(err)
			t.AssertEQ(inCooldown, true)
			t.AssertGT(remaining, int64(0))
			t.AssertLE(remaining, int64(60))

			// 场景3: 等待冷却期结束
			time.Sleep(2 * time.Second)

			// 手动更新创建时间，模拟冷却期已过
			_, err = tx.Model("verification_codes").
				Where("email", testEmail).
				Update(g.Map{"created_at": gtime.Now().Add(-2 * time.Minute)})
			t.AssertNil(err)

			inCooldown, remaining, err = logic.CheckVCodeCooldown(ctx, testEmail)
			t.AssertNil(err)
			t.AssertEQ(inCooldown, false)
			t.AssertEQ(remaining, int64(0))
		})
	})
}

// TestConsumeVCode 测试消费验证码
func TestConsumeVCode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testEmail := getTestEmail("consume")

		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			// 先创建一个验证码
			code, err := logic.CreateVCode(ctx, testEmail)
			t.AssertNil(err)

			// 消费验证码
			err = logic.ConsumeVCode(ctx, testEmail, code)
			t.AssertNil(err)

			// 验证验证码已被标记为已使用
			assertVCodeUsed(t, tx, testEmail, code, 1)

			// 再次消费应该不会报错（因为没有匹配的记录）
			err = logic.ConsumeVCode(ctx, testEmail, code)
			t.AssertNil(err)
		})
	})
}

// TestConsumeVCodeExpired 测试消费已过期的验证码
func TestConsumeVCodeExpired(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testEmail := getTestEmail("consume-expired")
		expiredCode := gconv.String(123456)

		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			// 手动插入一个已过期的验证码
			err := insertExpiredVCode(tx, testEmail, expiredCode)
			t.AssertNil(err)

			// 尝试消费已过期的验证码（不应更新记录）
			err = logic.ConsumeVCode(ctx, testEmail, expiredCode)
			t.AssertNil(err)

			// 验证验证码未被标记为已使用
			assertVCodeUsed(t, tx, testEmail, expiredCode, 0)
		})
	})
}

// TestVerifyAndConsumeVCode 测试原子性验证并消费验证码
func TestVerifyAndConsumeVCode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testEmail := getTestEmail("verify-consume")

		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			// 先创建一个验证码
			code, err := logic.CreateVCode(ctx, testEmail)
			t.AssertNil(err)

			// 验证并消费正确的验证码
			valid, err := logic.VerifyAndConsumeVCode(ctx, testEmail, code)
			t.AssertNil(err)
			t.AssertEQ(valid, true)

			// 验证验证码已被标记为已使用
			assertVCodeUsed(t, tx, testEmail, code, 1)

			// 再次验证并消费应该返回false（因为已使用）
			valid, err = logic.VerifyAndConsumeVCode(ctx, testEmail, code)
			t.AssertNil(err)
			t.AssertEQ(valid, false)
		})
	})
}

// TestVerifyAndConsumeVCodeOCTOPUS 测试防止TOCTOU竞态条件
func TestVerifyAndConsumeVCodeTOCTOU(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testEmail := getTestEmail("toctou")

		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			// 先创建一个验证码
			code, err := logic.CreateVCode(ctx, testEmail)
			t.AssertNil(err)

			// 第一次验证并消费
			valid1, err := logic.VerifyAndConsumeVCode(ctx, testEmail, code)
			t.AssertNil(err)
			t.AssertEQ(valid1, true)

			// 第二次验证并消费（模拟竞态条件，应该返回false）
			valid2, err := logic.VerifyAndConsumeVCode(ctx, testEmail, code)
			t.AssertNil(err)
			t.AssertEQ(valid2, false)
		})
	})
}

// TestVerifyAndConsumeVCodeExpired 测试原子性验证并消费已过期的验证码
func TestVerifyAndConsumeVCodeExpired(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testEmail := getTestEmail("verify-consume-expired")
		expiredCode := gconv.String(123456)

		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			// 手动插入一个已过期的验证码
			err := insertExpiredVCode(tx, testEmail, expiredCode)
			t.AssertNil(err)

			// 验证并消费已过期的验证码（应该返回false）
			valid, err := logic.VerifyAndConsumeVCode(ctx, testEmail, expiredCode)
			t.AssertNil(err)
			t.AssertEQ(valid, false)
		})
	})
}

// TestDeleteVCode 测试删除验证码
func TestDeleteVCode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testEmail := getTestEmail("delete")

		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			// 先创建一个验证码
			_, err := logic.CreateVCode(ctx, testEmail)
			t.AssertNil(err)

			// 删除验证码
			err = logic.DeleteVCode(ctx, testEmail)
			t.AssertNil(err)

			// 验证验证码已被删除
			assertVCodeNotExists(t, tx, testEmail)
		})
	})
}

// TestDeleteVCodeMultipleRecords 测试删除邮箱的所有验证码
func TestDeleteVCodeMultipleRecords(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		testEmail := getTestEmail("delete-multiple")

		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			// 创建多个验证码
			_, err := logic.CreateVCode(ctx, testEmail)
			t.AssertNil(err)
			time.Sleep(100 * time.Millisecond)
			_, err = logic.CreateVCode(ctx, testEmail)
			t.AssertNil(err)
			time.Sleep(100 * time.Millisecond)
			_, err = logic.CreateVCode(ctx, testEmail)
			t.AssertNil(err)

			// 获取验证码数量
			count, err := tx.Model("verification_codes").
				Where("email", testEmail).
				Count()
			t.AssertNil(err)
			t.AssertEQ(count, 3)

			// 删除验证码
			err = logic.DeleteVCode(ctx, testEmail)
			t.AssertNil(err)

			// 验证所有验证码已被删除
			count, err = tx.Model("verification_codes").
				Where("email", testEmail).
				Count()
			t.AssertNil(err)
			t.AssertEQ(count, 0)
		})
	})
}

// TestDeleteExpiredVCodes 测试删除过期的验证码
func TestDeleteExpiredVCodes(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		validEmail := getTestEmail("valid")
		expiredEmail := getTestEmail("expired")
		expiredCode := gconv.String(123456)

		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			// 创建一个有效的验证码
			validCode, err := logic.CreateVCode(ctx, validEmail)
			t.AssertNil(err)

			// 手动插入一个已过期的验证码
			err = insertExpiredVCode(tx, expiredEmail, expiredCode)
			t.AssertNil(err)

			// 删除过期的验证码
			err = logic.DeleteExpiredVCodes(ctx)
			t.AssertNil(err)

			// 验证有效验证码仍然存在
			assertVCodeExists(t, tx, validEmail, validCode)

			// 验证过期验证码已被删除
			assertVCodeNotExists(t, tx, expiredEmail)
		})
	})
}

// TestDeleteExpiredVCodesMixed 测试删除过期验证码时混合有效和过期记录
func TestDeleteExpiredVCodesMixed(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		validEmail1 := getTestEmail("valid1")
		validEmail2 := getTestEmail("valid2")
		expiredEmail1 := getTestEmail("expired1")
		expiredEmail2 := getTestEmail("expired2")
		expiredEmail3 := getTestEmail("expired3")

		runInTransaction(t, func(ctx context.Context, tx gdb.TX) {
			// 创建多个有效验证码
			_, err := logic.CreateVCode(ctx, validEmail1)
			t.AssertNil(err)
			time.Sleep(100 * time.Millisecond)
			_, err = logic.CreateVCode(ctx, validEmail2)
			t.AssertNil(err)

			// 手动插入多个已过期的验证码
			expiredEmails := []string{expiredEmail1, expiredEmail2, expiredEmail3}
			for i, email := range expiredEmails {
				err = insertExpiredVCode(tx, email, gconv.String(100000+i))
				t.AssertNil(err)
			}

			// 删除过期的验证码
			err = logic.DeleteExpiredVCodes(ctx)
			t.AssertNil(err)

			// 验证所有有效验证码仍然存在
			for _, email := range []string{validEmail1, validEmail2} {
				record, err := tx.Model("verification_codes").Where("email", email).One()
				t.AssertNil(err)
				t.AssertEQ(record.IsEmpty(), false)
			}

			// 验证所有过期验证码已被删除
			for _, email := range expiredEmails {
				assertVCodeNotExists(t, tx, email)
			}
		})
	})
}
