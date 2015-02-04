package coderg

import (
	"regexp"
	"strings"
	"strconv"
)

/*
 * InputProcessor
 * 输入处理器
 */


type InputProcessor struct {
	reScript  *regexp.Regexp;
	reMark    *regexp.Regexp;
	reEmail   *regexp.Regexp;
	reUrl   *regexp.Regexp;
}

func NewInputProcessor() (ip *InputProcessor){
	ip = new(InputProcessor);
	ip.reScript, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>");
	ip.reMark, _   = regexp.Compile("^[A-Za-z0-9_-]+$");
	ip.reEmail, _  = regexp.Compile("^[A-Za-z0-9]+([_.-][A-Za-z0-9]+)*@[A-Za-z0-9]+([_.-][A-Za-z0-9]+)*.([A-Za-z]){2,5}$");
	ip.reUrl, _  = regexp.Compile("^[A-Za-z0-9]+://");
	return;
}

/*
 * replaceText
 * 替换危险字符
 */
func (ip *InputProcessor) replaceText(text string) string{
	theReplaceMap := make(map[string]string);
	theReplaceMap["<"] = "&lt;";
	theReplaceMap[">"] = "&gt;";
	theReplaceMap["\""] = "&guot;";
	theReplaceMap["'"] = "&#039;";
	theReplaceMap["|"] = "&brvbar;";
	theReplaceMap["`"] = "&acute;";
	for index, value := range theReplaceMap {
		text = strings.Replace(text, index, value, -1);
	}
	return text;
}


/*
 * minOrMax
 * 按要求判断字符长度是否符合要求
 */
func (ip *InputProcessor) minOrMax(text string, min, max int) (err int){
	textLen := len(text);
	if(textLen < min){
		err = 2;
		return;
	}
	if(max != 0 && textLen > max){
		err = 3;
		return;
	}
	err = 0;
	return;
}

/*
 * Text
 * 处理简单文本输入，替换可能引起注入的字符、判断长度等
 */
func (ip *InputProcessor) Text(text string, must bool, min, max int) (textc string, err int){
	err = 0;
	textc = "";
	text = strings.TrimSpace(text);
	textLen := len(text);
	if(must == true && textLen == 0){
		err = 1;
		return;
	}
	if(textLen == 0){
		return;
	}
	text = ip.replaceText(text);
	err = ip.minOrMax(text,min,max);
	if err == 0	{
		textc = text;
	}
	return;
}

/*
 * Mark
 * 处理作为标记的的字符串，只能由字母数字和连字符组成
 */
func (ip *InputProcessor) Mark(text string, must bool, min, max int) (textc string, err int){
	err = 0;
	textc = "";
	text = strings.TrimSpace(text);
	textLen := len(text);
	if(must == true && textLen == 0){
		err = 1;
		return;
	}
	if(textLen == 0){
		return;
	}
	err = ip.minOrMax(text,min,max);
	if err != 0	{
		return;
	}
	ifmatch := ip.reMark.MatchString(text)
	if ( ifmatch == false){
		err = 4;
		return;
	}
	textc = text;
	return;
}

/*
 * EditorIn
 * 处理编辑器的输入，除了字数检测外，主要是过滤script标签
 */
func (ip *InputProcessor) EditorIn(text string, must bool, min, max int) (textc string, err int){
	err = 0;
	textc = "";
	text = strings.TrimSpace(text);
	textLen := len(text);
	if(must == true && textLen == 0){
		err = 1;
		return;
	}
	if(textLen == 0){
		return;
	}
	
    text = ip.reScript.ReplaceAllString(text, "");
    text = strings.Replace(text, "`", "&acute;", -1);
    text = strings.Replace(text, "'", "''", -1);
    err = ip.minOrMax(text,min,max);
	if err != 0	{
		return;
	}
	textc = text;
	return;
}

/*
 * EditorRe
 * 处理编辑器的重新编辑
 */
func (ip *InputProcessor) EditorRe(text string) (string){
	text = strings.Replace(text, "&acute;", "`", -1);
    text = strings.Replace(text, "&#039;", "''", -1);
    return text;
}

/*
 * TextareaOut
 * 处理文本域的输出，主要就是对换行符进行处理，将\r\n之类的转换成<p>或<br>
 * thetype为true则转成<p>，type为false则转为<br>
 * 
 */
func (ip *InputProcessor) TextareaOut (text string, thetype bool) string{
	if (thetype == true) {
		theReplaceMap := make(map[string]string);
		theReplaceMap["\r\n"] = "</p><p>";
		theReplaceMap["\r"] = "</p><p>";
		theReplaceMap["\n"] = "</p><p>";
		for index, value := range theReplaceMap {
			text = strings.Replace(text, index, value, -1);
		}
		text = "<p>" + text + "</p>";
	}else{
		theReplaceMap := make(map[string]string);
		theReplaceMap["\r\n"] = "<br>";
		theReplaceMap["\r"] = "</br>";
		theReplaceMap["\n"] = "<br>";
		for index, value := range theReplaceMap {
			text = strings.Replace(text, index, value, -1);
		}
	}
	return text;
}

/*
 * Int
 * 输入是否为整数
 * 
 */
func (ip *InputProcessor) Int(text string, must bool, min, max int64) (num int64, err int){
	text = strings.TrimSpace(text);
	err = 0;
	textLen := len(text);
	if(must == true && textLen == 0){
		num = 0;
		err = 1;
		return;
	}
	if(textLen == 0){
		num = 0;
		return;
	}
	num, e := strconv.ParseInt(text,10,64);
	if e != nil {
		err = 2;
		return;
	}
	if (num < min || num > max){
		err = 3;
		return;
	}
	return;
}

/*
 * Float
 * 输入是否为浮点
 * 
 */
func (ip *InputProcessor) Float (text string, must bool, min, max float64) (num float64, err int){
	text = strings.TrimSpace(text);
	err = 0;
	textLen := len(text);
	if(must == true && textLen == 0){
		num = 0;
		err = 1;
		return;
	}
	if(textLen == 0){
		num = 0;
		return;
	}
	num, e := strconv.ParseFloat(text,64);
	if e != nil {
		err = 2;
		return;
	}
	if (num < min || num > max){
		err = 3;
		return;
	}
	return;
}

/*
 * Enum
 * 枚举，判断提供的字符串是否出现在字符串切片里
 */
func (ip *InputProcessor) Enum(text string, must bool, enum []string) (textc string, err int){
	err = 0;
	textc = "";
	text = strings.TrimSpace(text);
	textLen := len(text);
	if(must == true && textLen == 0){
		err = 1;
		return;
	}
	if(textLen == 0){
		return;
	}
	for _, v := range enum {
		if text == v {
			textc = text;
			return;
		}
	}
	err = 2;
	return;
}

/*
 * Email
 * 查看格式是不是邮箱地址
 */
func (ip *InputProcessor) Email(text string, must bool, min, max int)(textc string, err int){
	err = 0;
	textc = "";
	text = strings.TrimSpace(text);
	textLen := len(text);
	if(must == true && textLen == 0){
		err = 1;
		return;
	}
	if(textLen == 0){
		return;
	}
	err = ip.minOrMax(text,min,max);
	if err != 0	{
		return;
	}
	ifmatch := ip.reEmail.MatchString(text)
	if ( ifmatch == false){
		err = 4;
		return;
	}
	textc = text;
	return;
}

/*
 * Url
 * 处理连接，主要是如果没有协议的话，就加上http://
 */
func (ip *InputProcessor) Url (text string, must bool, min, max int) (textc string, err int) {
	err = 0;
	textc = "";
	text = strings.TrimSpace(text);
	textLen := len(text);
	if(must == true && textLen == 0){
		err = 1;
		return;
	}
	if(textLen == 0){
		return;
	}
	err = ip.minOrMax(text,min,max);
	if err != 0	{
		return;
	}
	ifmatch := ip.reUrl.MatchString(text)
	if ( ifmatch == false){
		textc = "http://" + text;
	}else{
		textc = text;
	}
	return;
}

/*
 * Regular
 * 正则判断，提供一个正则表达，如果匹配则返回字符串
 */
func (ip *InputProcessor) Regular (text string, must bool ,rg string) (textc string, err int){
	err = 0;
	textc = "";
	text = strings.TrimSpace(text);
	textLen := len(text);
	if(must == true && textLen == 0){
		err = 1;
		return;
	}
	if(textLen == 0){
		return;
	}
	therp, ther := regexp.MatchString(rg, text);
	if (ther != nil){
		err = 2;
		return;
	}
	if (therp == false) {
		err = 3;
		return;
	}
	textc = text;
	return;
}
