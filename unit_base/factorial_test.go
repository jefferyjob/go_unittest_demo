package unit_base

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	// 程序输出的结果
	got := Factorial(5)

	// 期望的结果
	want := &Fact{
		ret:  120,
		nums: []int{5, 4, 3, 2, 1},
	}

	// 因为struct不能比较直接，借助反射包中的方法比较
	if !reflect.DeepEqual(want, got) {
		// 测试失败输出错误提示
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// 这里给到一个错误的结果，测试错误情况
func TestFactEr(t *testing.T) {
	got := Factorial(3)
	want := &Fact{
		ret:  6,
		nums: []int{3, 2, 1},
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// 子测试演示
func TestFactGroup(t *testing.T) {
	t.Parallel() // 将 TLog 标记为能够与其他测试并行运行

	// 定义测试表格
	testCases := []struct {
		testName string
		input    int
		output   *Fact
	}{
		{"case1", 5, &Fact{ret: 120, nums: []int{5, 4, 3, 2, 1}}},
		{"case2", 3, &Fact{ret: 6, nums: []int{3, 2, 1}}},
	}

	// 运行子测试代码
	for _, tt := range testCases {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel() // 将每个测试用例标记为能够彼此并行运行
			got := Factorial(tt.input)
			if !reflect.DeepEqual(tt.output, got) {
				t.Errorf("expected:%v, got:%v", tt.output, got)
			}
		})
	}
}

// 基准测试
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(10)
	}
}

// 基准测试之性能比较测试
func benchmarkFact(b *testing.B, x int) {
	for i := 0; i < b.N; i++ {
		Factorial(x)
	}
}
func BenchmarkFact1(b *testing.B)   { benchmarkFact(b, 1) }
func BenchmarkFact5(b *testing.B)   { benchmarkFact(b, 5) }
func BenchmarkFact10(b *testing.B)  { benchmarkFact(b, 10) }
func BenchmarkFact50(b *testing.B)  { benchmarkFact(b, 50) }
func BenchmarkFact100(b *testing.B) { benchmarkFact(b, 100) }

// 基准测试之并行测试
func BenchmarkFactorialParallel(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Factorial(10)
		}
	})
}

// 测试集的Setup与Teardown
func setupFactsTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:Setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:Setup")
	}
}

// 子测试的Setup与Teardown
func downFactsTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:Teardown")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:Teardown")
	}
}

func TestFactsGroup(t *testing.T) {
	t.Parallel()

	// 定义测试表格
	testCases := []struct {
		testName string
		input    int
		output   *Fact
	}{
		{"case1", 5, &Fact{ret: 120, nums: []int{5, 4, 3, 2, 1}}},
		{"case2", 3, &Fact{ret: 6, nums: []int{3, 2, 1}}},
	}

	setupTest := setupFactsTest(t)
	defer setupTest(t) // 测试之前执行setup操作

	// 运行子测试代码
	for _, tt := range testCases {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()

			downTest := downFactsTest(t)
			defer downTest(t) // 测试之后执行testdown操作

			got := Factorial(tt.input)
			if !reflect.DeepEqual(tt.output, got) {
				t.Errorf("expected:%v, got:%v", tt.output, got)
			}
		})
	}
}

// Example
func ExampleFactorial() {
	fmt.Println(Factorial(5))

	// Output:&Fact{ret:120, nums:[]int{5,4,3,2,1}}
}

// gotests 工具生成的测试代码
//func Test_operation(t *testing.T) {
//	type args struct {
//		x     int
//		factS *Fact
//	}
//	tests := []struct {
//		name string
//		args args
//		want *Fact
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := operation(tt.args.x, tt.args.factS); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("operation() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
