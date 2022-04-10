package g

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Number", func() {
	It("ForceInt should work", func() {
		var s = "123"
		Expect(ForceInt(s, 0)).To(Equal(123))
		s = "abc"
		Expect(ForceInt(s, -1)).To(Equal(-1))
	})

	It("ForceInt32 should work", func() {
		var s = "123"
		Expect(ForceInt32(s, 0)).To(Equal(int32(123)))
		s = "abc"
		Expect(ForceInt32(s, -1)).To(Equal(int32(-1)))
	})

	It("ForceInt64 should work", func() {
		var s = "123"
		Expect(ForceInt64(s, 0)).To(Equal(int64(123)))
		s = "abc"
		Expect(ForceInt64(s, -1)).To(Equal(int64(-1)))
	})

	It("String should work", func() {
		var s string = "123456"

		Expect(String(s)).To(Equal("123456"))
		var i int = 10086

		Expect(String(i)).To(Equal("10086"))

		var i32 int32 = 12345
		Expect(String(i32)).To(Equal("12345"))

		var i64 int64 = 1258096
		Expect(String(i64)).To(Equal("1258096"))

		var f32 float32 = 123.321
		Expect(String(f32, 3)).To(Equal("123.321"))

		var in float64 = 1234567.7654321
		Expect(String(in, 7)).To(Equal("1234567.7654321"))

	})

	It("Int should work", func() {
		var s string = "123"
		result, _ := Int(s)
		Expect(result, int(123))

		var i int = 10086
		result, _ = Int(i)
		Expect(result).To(Equal(i))

		var i32 int32 = 12345
		result, _ = Int(i32)
		Expect(result, int(i32))

		var i64 int64 = 1258096
		result, _ = Int(i64)
		Expect(result, int(i64))

		var f32 float32 = 123.321
		result, _ = Int(f32)
		Expect(result, int(f32))

		var f64 float64 = 1234567.7654321
		result, _ = Int(f64)
		Expect(result, int(f64))
	})

	It("Int32 should work", func() {
		var s string = "123"
		result, _ := Int32(s)
		Expect(result, int32(123))

		var i int = 10086
		result, _ = Int32(i)
		Expect(result).To(Equal(int32(i)))

		var i32 int32 = 12345
		result, _ = Int32(i32)
		Expect(result).To(Equal(i32))

		var i64 int64 = 1258096
		result, _ = Int32(i64)
		Expect(result).To(Equal(int32(i64)))

		var f32 float32 = 123.321
		result, _ = Int32(f32)
		Expect(result).To(Equal(int32(f32)))

		var f64 float64 = 1234567.7654321
		result, _ = Int32(f64)
		Expect(result).To(Equal(int32(f64)))
	})

	It("Int64 should work", func() {
		var s string = "123"
		result, _ := Int64(s)
		Expect(result).To(Equal(int64(123)))

		var i int = 10086
		result, _ = Int64(i)
		Expect(result).To(Equal(int64(i)))

		var i32 int32 = 12345
		result, _ = Int64(i32)
		Expect(result).To(Equal(int64(i32)))

		var i64 int64 = 1258096
		result, _ = Int64(i64)
		Expect(result).To(Equal(i64))

		var f32 float32 = 123.321
		result, _ = Int64(f32)
		Expect(result).To(Equal(int64(f32)))

		var f64 float64 = 1234567.7654321
		result, _ = Int64(f64)
		Expect(result).To(Equal(int64(f64)))
	})

	It("Float32 should work", func() {
		var s string = "123.456"
		result, _ := Float32(s)
		Expect(result).To(Equal(float32(123.456)))

		var i int = 10086
		result, _ = Float32(i)
		Expect(result).To(Equal(float32(i)))

		var i32 int32 = 12345
		result, _ = Float32(i32)
		Expect(result).To(Equal(float32(i32)))

		var i64 int64 = 1258096
		result, _ = Float32(i64)
		Expect(result).To(Equal(float32(i64)))

		var f32 float32 = 123.321
		result, _ = Float32(f32)
		Expect(result).To(Equal(f32))

		var f64 float64 = 1234567.7654321
		result, _ = Float32(f64)
		Expect(result, float32(f64))
	})

	It("Float64 should work", func() {
		var s string = "123.456"
		result, _ := Float64(s)
		Expect(result).To(Equal(float64(123.456)))

		var i int = 10086
		result, _ = Float64(i)
		Expect(result, float64(i))

		var i32 int32 = 12345
		result, _ = Float64(i32)
		Expect(result, float64(i32))

		var i64 int64 = 1258096
		result, _ = Float64(i64)
		Expect(result, float64(i64))

		var f32 float32 = 123.321
		result, _ = Float64(f32)
		Expect(result, float64(f32))

		var f64 float64 = 1234567.7654321
		result, _ = Float64(f64)
		Expect(result).To(Equal(f64))
	})
})
