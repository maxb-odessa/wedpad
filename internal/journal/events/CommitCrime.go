package events

// CommitCrime event structure
type CommitCrimeT interface{}

// CommitCrime event handler
func (evHandler EventHandler) CommitCrime(eventData map[string]interface{}) {
    // ev := new(CommitCrimeT)
    // mapstructure.Decode(eventData, ev)
}

