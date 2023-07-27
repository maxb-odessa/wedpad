
# quick and dirty script to generate Go event structs from ED journals (json to go)

LIST=$(cat *.log *.json Journals-old/*.log | jq .event | sort | uniq | tr -d '"')

mkdir -p events

for EV in $LIST; do

    cat Journal*log | \
    grep "\"event\":\"$EV\"" | \
    gojsonstruct -typename ${EV}T -packagename events -structtagname mapstructure -omitempty true -typecomment "$EV event structure" \
    > events/$EV.go

    if [ -s $EV.json ]; then
        cat $EV.json | \
        gojsonstruct -typename ${EV}T -packagename events -structtagname mapstructure -omitempty true -typecomment "$EV event structure" \
        > events/$EV.go
    fi

cat<<EOF>> events/$EV.go

// $EV event handler
func (evh *EventHandler) $EV(eventData map[string]interface{}) {
    // ev := new(${EV}T)
    // mapstructure.Decode(eventData, ev)
    // cs := evh.CurrentSystem()
}

EOF

done

