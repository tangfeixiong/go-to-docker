package dispatcher

type AgentManager struct {
	CollectorDriver string
	Duration        int64
}

func (am *AgentManager) Dispatch(chsig <-chan bool) {
	<-chsig
}
