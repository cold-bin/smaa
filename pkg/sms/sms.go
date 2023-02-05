// @author cold bin
// @date 2023/2/5

package sms

import (
	"eliminate-male-appearance-anxiety/global"
	"errors"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	_sms "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/cold-bin/goutil/general/mlog"
)

func InitSmsClient() (err error) {
	smsConf := global.Set.Sms
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: tea.String(smsConf.AccessKeyId),
		// 必填，您的 AccessKey Secret
		AccessKeySecret: tea.String(smsConf.AccessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	global.SmsClient = &_sms.Client{}
	global.SmsClient, err = _sms.NewClient(config)

	return err
}

// aJudgeStatus 简单判断阿里云的请求是否成功。如果没成功，直接返回错误消息
func aJudgeStatus(status string) (err error) {
	if status == "" {
		return errors.New("the status is empty")
	} else if status == "OK" {
		return nil
	}
	return errors.New(fmt.Sprintf("aliyun responce err code: %s", status))
}

func Send(phone string, code string) (err error) {
	smsConf := global.Set.Sms
	sendSmsRequest := &_sms.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(smsConf.SignName),
		TemplateCode:  tea.String(smsConf.TemplateCode),
		TemplateParam: tea.String(fmt.Sprintf(smsConf.TemplateParam, code)),
	}

	runtime := &util.RuntimeOptions{}
	tryErr := func() (e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		r, err := global.SmsClient.SendSmsWithOptions(sendSmsRequest, runtime)
		// todo 第三方平台的状态码处理
		if err != nil {
			return err
		}

		// 处理阿里云业务状态码
		if err = aJudgeStatus(*r.Body.Code); err != nil {
			return err
		}

		mlog.Debugf("短信发送：", r.String())
		return err
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		_, err = util.AssertAsString(error.Message)
		if err != nil {
			return
		}
	}
	return
}
