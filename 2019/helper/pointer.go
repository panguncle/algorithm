package helper

/**
IntPtr: 获得i的指针, 由于golang是值传递, 因此 i 传进来地址已经与外部的地址不同
因此获得指针的地址也与外部的不同
*/
func IntPtr(i int) *int {
	return &i
}
