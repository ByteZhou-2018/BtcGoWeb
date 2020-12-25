package mycrypts

import "encoding/base64"
/**
*基于可以打印字符来表示二进制数据的编码方式。base64使用了26个大写字母、26六个小写字母、十个数字和 + /
* 			base64编码通常被用于编码邮件中的附件
*将每个字符转成ASCII编码（十进制）
*将十进制编码转换成二级制编码
*将二级制编码按照六位一组进行平分
*将六位一组的二进制数高位补零，然后转成十进制数
*以十进制数作为索引，去base64编码表中去查找字符
*每三个字符的文本将编码为四个字符长度。
				若文本为3个字符，则正好编码为4个字符长度
				若文本为2个字符，由于不足4个字符，则在尾部用 = 补齐
				若文本为1个字符，则编码为2个字符，则在尾部用两个 = 补齐。
 */
func Base64EncodeString(str string)string  {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
func Base64DecodeString(str string)string  {
	bytes, _ := base64.StdEncoding.DecodeString(str)
	return string(bytes)
}
