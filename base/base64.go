package base

import (
	"encoding/base64"
)

/**
 * base64编码
 */
func Base64Encode(data []byte) []byte{
	/**
	 * base64: 64个字符，A-Z，a-z，0-9，+ /
	 * 3 * 8 = 24
	 * ？ * 6 = 24
	 * 分组后每一组6位：111111 -> 63、000000 -> 0
	 * 一共64：
	 * 原理：
	 *   ① 数据 -> ASCII编码（十进制）
	 *   ② 十进制 -> 二进制（0和1的组合）
	 *   ③ 6位一组进行分组
	 *   ④ 将6位一组的二进制转换为十进制（0-63）
	 *   ⑤ 将转换后的十进制数作为元素的下标到固定的编码表中获取对应位置的元素


	 * 将原数据进行24位一组分组，分组情况有三种：
		① 刚刚满足24位一组的分组，恰好
		② 余数为1个字符，最后补两个=
		③ 余数为2个字符，最后补一个=
	 */

	encoding := base64.StdEncoding
	dst := make([]byte, 0)
	encoding.Encode(dst,data)
	return dst
}

/**
 * base64编码
 */
func Base64Decode(data string) []byte{
	encoding := base64.StdEncoding
	dst := make([]byte, 0)
	encoding.Decode(dst, []byte(data))
	return dst
}
