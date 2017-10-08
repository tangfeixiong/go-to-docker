package dispatcher

type CollectorManager struct {
	StorageDriver string
	Duration      int64
}

func (cm *CollectorManager) Dispatch(chsig <-chan bool) {
	<-chsig
}
