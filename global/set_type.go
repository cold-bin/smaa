// @author cold bin
// @date 2023/2/4

package global

// Settings 项目配置
type Settings struct {
	App   App   `mapstructure:"app"`
	Mysql Mysql `mapstructure:"mysql"`
	Redis Redis `mapstructure:"redis"`
	Sms   Sms   `mapstructure:"sms"`
	Oss   Oss   `mapstructure:"oss"`
}

type Oss struct {
	AccessKey      string `mapstructure:"access_key"`
	SecretKey      string `mapstructure:"secret_key"`
	Bucket         string `mapstructure:"bucket"`
	UploadTokenExp uint64 `mapstructure:"upload_token_exp"`
	Domain         string `mapstructure:"domain"`
	ImageRoot      string `mapstructure:"image_root"`
	Video          string `mapstructure:"video_root"`
}
type Sms struct {
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	SignName        string `mapstructure:"sign_name"`
	TemplateCode    string `mapstructure:"template_code"`
	TemplateParam   string `mapstructure:"template_param"`
}

type App struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
}

type Mysql struct {
	UserName        string `mapstructure:"user_name"`
	Password        string `mapstructure:"password"`
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	Db              string `mapstructure:"db"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}
