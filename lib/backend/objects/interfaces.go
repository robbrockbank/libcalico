package objects

type Key interface {
	asEtcdKey() string
}