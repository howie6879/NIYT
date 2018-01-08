package common

import (
	"math/rand"
	"time"
)

// UserAgents a slice contains ua
var (
	UserAgents = []string{
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; Media Center PC 6.0; InfoPath.3; MS-RTC LM 8; Zune 4.7)",
		"Mozilla/4.0 (compatible; MSIE 9.0; Windows NT 5.1; Trident/5.0)",
		"Mozilla/5.0 (X11; Linux x86_64; rv:2.2a1pre) Gecko/20100101 Firefox/4.2a1pre",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:2.0b11pre) Gecko/20110131 Firefox/4.0b11pre",
		"Mozilla/5.0 (X11; U; Linux i686; ru-RU; rv:1.9.2a1pre) Gecko/20090405 Ubuntu/9.04 (jaunty) Firefox/3.6a1pre",
		"Mozilla/5.0 (X11; U; Linux x86_64; en-US; rv:1.9.2.8) Gecko/20100723 SUSE/3.6.8-0.1.1 Firefox/3.6.8",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; pt-PT; rv:1.9.2.6) Gecko/20100625 Firefox/3.6.6",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; it; rv:1.9.2.6) Gecko/20100625 Firefox/3.6.6 ( .NET CLR 3.5.30729)",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.2.6) Gecko/20100625 Firefox/3.6.6 (.NET CLR 3.5.30729)",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; ru; rv:1.9.2.4) Gecko/20100513 Firefox/3.6.4",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; ja; rv:1.9.2.4) Gecko/20100611 Firefox/3.6.4 GTB7.1",
	}
	DomainFlagM       = "www.so.com/link?m"
	DomainFlagL       = "www.so.com/link?url"
	SoURL             = "https://www.so.com/s"
	LatestChapterName = "og:novel:latest_chapter_name"
	LatestChapterURL  = "og:novel:latest_chapter_url"
	Sites             = []string{
		"www.58xs.tw",
		"www.nuomi9.com",
		"www.xiaoshuoli.com",
		"www.biquguan.com",
		"www.23us.la",
		"www.xs98.com",
		"www.biqugex.com",
		"www.biquge.tw",
		"www.qu.la",
		"www.ybdu.com",
		"www.wenxuemi.com",
		"www.biquge.com",
		"www.23us.cc",
		"www.xs222.com",
		"www.lewen8.com",
		"www.bqg5200.com",
		"www.vodtw.com",
		"www.6mao.com",
		"www.touxiang.la",
		"www.bxquge.com",
		"www.beidouxin.com",
		"www.263zw.com",
		"www.3qzone.com",
		"wwww.yooread.com",
		"www.suimeng.la",
		"www.bequge.com",
		"www.biquku.co",
		"www.xbqge.com",
		"www.aiquxs.com",
		"www.23us.com",
		"www.ddbiquge.com",
		"www.abocms.cn",
		"www.liewen.cc",
		"www.8535.org",
		"www.dingdianzw.com",
		"www.biquge.cc",
		"www.111bz.org",
		"www.biqugebook.com",
		"www.e8zw.com",
		"www.xqqxs.com",
		"tianyibook.la",
		"www.lingdianksw.com",
		"www.qb5.tw",
		"www.quanben.com",
		"www.58xs.com",
		"www.biqukan.com",
		"www.yssm.org",
		"www.81zw.com",
		"www.ymoxuan.com",
		"www.mytxt.cc",
		"www.woquge.com",
		"www.biquguo.com",
		"www.8jzw.cc",
		"www.8jzw.com",
		"www.23xsw.cc",
		"www.miaobige.com",
		"www.xs.la",
		"www.44pq.co",
		"www.50zw.la",
		"www.33xs.com",
		"www.zwdu.com",
		"www.ttzw.com",
		"www.biqudu.com",
		"www.biqugeg.com",
		"www.23txt.com",
		"www.baquge.tw",
		"www.lread.cc",
		"www.laidudu.com",
		"www.kxs7.com",
		"www.biquguan.com",
		"www.biquta.com",
		"www.xs98.com",
		"www.bqge.org",
		"www.58xs.tw",
		"www.187ks.com",
		"www.yikanxiaoshuo.com",
		"www.23zw.me",
		"www.37zw.net",
		"www.biquge.cm",
		"www.kanshu58.com",
		"www.biqumo.com",
		"www.mpxiaoshuo.com",
		"www.23wx.cm",
		"www.biquge.jp",
		"www.biqugexsw.com",
		"www.biqu6.com",
		"www.xiuxs.com",
		"www.biqule.com",
		"www.biquzi.com",
		"www.biquku.la",
		"www.00ksw.org",
		"www.bqg.cc",
		"www.biqugezw.com",
	}
)

// LOGO show the basic info
const LOGO = `
███╗   ██╗██╗██╗   ██╗████████╗
████╗  ██║██║╚██╗ ██╔╝╚══██╔══╝
██╔██╗ ██║██║ ╚████╔╝    ██║   
██║╚██╗██║██║  ╚██╔╝     ██║   
██║ ╚████║██║   ██║      ██║   
╚═╝  ╚═══╝╚═╝   ╚═╝      ╚═╝   

Read the novel in your terminal - NIYT v0.1.2

`

// Config contains information we need to process a novel
type Config struct {
	DomainFlagM    string
	DomainFlagL    string
	SoURL          string
	LatestChapeter struct {
		LatestChapterName string
		LatestChapterURL  string
	}
	Sites []string
}

// LoadConfiguration Get the configuration from the rules.json
func LoadConfiguration() Config {
	var config Config
	config.DomainFlagM = DomainFlagM
	config.DomainFlagL = DomainFlagL
	config.SoURL = SoURL
	config.LatestChapeter.LatestChapterName = LatestChapterName
	config.LatestChapeter.LatestChapterURL = LatestChapterURL
	config.Sites = Sites
	return config
}

// GetUserAgent Get a random user agent string.
func GetUserAgent() string {
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(UserAgents)
	return UserAgents[n]
}
