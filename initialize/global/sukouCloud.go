package global

var (
	Spec            string //定时任务方法
	BoxProject      map[string]map[int]int
	SukonCloudToken string //全局变量,速控云的token
	SendBoxTaskChan chan string
)

func SukonCloudChan() {
	SendBoxTaskChan = make(chan string, 1)
}
