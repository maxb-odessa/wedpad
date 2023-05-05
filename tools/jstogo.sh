
# quick and dirty script to generate Go event structs from ED journals (json to go)

LIST=$(cat *.log *.json | jq .event | sort | uniq | tr -d '"')

cat<<EOF>events/events.go
package events

type eventHandlersFunc func(map[string]interface{})

var eventHandlers map[string]eventHandlersFunc

func init() {
    eventHandlers = make(map[string]eventHandlersFunc)


EOF

for EV in $LIST; do

    cat Journal*log | \
    grep "\"event\":\"$EV\"" | \
    gojsonstruct -typename ev${EV} -packagename events -structtagname mapstructure -omitempty true -typecomment "$EV event structure" \
    > events/$EV.go

    if [ -s $EV.json ]; then
        cat $EV.json | \
        gojsonstruct -typename ev${EV} -packagename events -structtagname mapstructure -omitempty true -typecomment "$EV event structure" \
        > events/$EV.go
    fi

cat<<EOF>> events/$EV.go

// $EV event handler
func $EV(eventData map[string]interface{}) {
    // ev := new(ev$EV)
    // mapstructure.Decode(eventData, ev)
}

EOF

cat<<EOF>>events/events.go
    eventHandlers["$EV"] = $EV
EOF


done

cat<<EOF>>events/events.go
}

func Handle(event string, data map[string]interface{}) {
    if handler, ok := eventHandlers[event]; ok {
        slog.Debug(9, "handling event '%s'", event)
        handler(data)
    } else {
        slog.Err("Undefined event '%s', can't handle", event)
    }
}

EOF
