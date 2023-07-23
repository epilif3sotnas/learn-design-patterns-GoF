# std
import
    std/[
        json
    ]

# internal
import
    patterns/structural/adapter/[
        governmentadapter
    ]


type
    Broker*[TAdapter: GovernmentAdapter] = ref object
        adapter: TAdapter


proc newBroker*[TAdapter: GovernmentAdapter](adapter: TAdapter): Broker[TAdapter] =
    return Broker[TAdapter](
        adapter: adapter
    )

proc sendGains*(self: Broker, data: JsonNode): bool =
    return self.adapter.sendGainsInfo(data)