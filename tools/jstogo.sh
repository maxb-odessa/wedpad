
# quick and dirty script to generate Go event structs from ED journals (json to go)

LIST=$(cat *.log *.json | jq .event | sort | uniq | tr -d '"')

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
func (evHandler EventHandler) $EV(eventData map[string]interface{}) {
    // ev := new(ev$EV)
    // mapstructure.Decode(eventData, ev)
}

EOF

done

