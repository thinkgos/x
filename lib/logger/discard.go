package logger

// Discard is an logger on which all Write calls succeed
// without doing anything.
type Discard struct{}

var _ Logger = (*Discard)(nil)

// NewDiscard a discard logger on which always succeed without doing anything
func NewDiscard() Discard { return Discard{} }

// Debugf implement Logger interface.
func (sf Discard) Debugf(string, ...interface{}) {}

// Infof implement Logger interface.
func (sf Discard) Infof(string, ...interface{}) {}

// Errorf implement Logger interface.
func (sf Discard) Errorf(string, ...interface{}) {}

// Warnf implement Logger interface.
func (sf Discard) Warnf(string, ...interface{}) {}

// DPanicf implement Logger interface.
func (sf Discard) DPanicf(string, ...interface{}) {}

// Fatalf implement Logger interface.
func (sf Discard) Fatalf(string, ...interface{}) {}
