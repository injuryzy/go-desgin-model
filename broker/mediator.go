package broker

type Mediator interface {
	CanArrive(train Train) bool
	NotifyAboutDepart()
}

type StaticMediator struct {
	IsPlantFrom bool
	TrainQueue  []Train
}

func NewStaticMediator() *StaticMediator {
	return &StaticMediator{
		IsPlantFrom: true,
	}
}

func (s *StaticMediator) CanArrive(train Train) bool {
	if s.IsPlantFrom {
		s.IsPlantFrom = false
		return true
	}
	s.TrainQueue = append(s.TrainQueue, train)
	return false
}

func (s *StaticMediator) NotifyAboutDepart() {
	if !s.IsPlantFrom {
		s.IsPlantFrom = true
	}
	if len(s.TrainQueue) > 0 {
		firstTrain := s.TrainQueue[0]
		firstTrain.CanArrive()
		s.TrainQueue = s.TrainQueue[1:]
	}
}
