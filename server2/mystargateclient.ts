import { fromUtf8 } from "@cosmjs/encoding"
import {
    Attribute as TendermintAttribute,
    BlockResultsResponse,
    Event,
    Tendermint34Client,
} from "@cosmjs/tendermint-rpc"
import { StargateClient } from "@cosmjs/stargate"
import {
    Attribute,
    StringEvent,
} from "cosmjs-types/cosmos/base/abci/v1beta1/abci"

export class MyStargateClient extends StargateClient {
    private readonly myTmClient: Tendermint34Client

    public static async connect(endpoint: string): Promise<MyStargateClient> {
        const tmClient = await Tendermint34Client.connect(endpoint)
        return new MyStargateClient(tmClient)
    }

    protected constructor(tmClient: Tendermint34Client) {
        super(tmClient)
        this.myTmClient = tmClient
    }

    protected convertTendermintEvents(events: readonly Event[]): StringEvent[] {
        return events.map(
            (event: Event): StringEvent => ({
                type: event.type,
                attributes: event.attributes.map(
                    (attribute: TendermintAttribute): Attribute => ({
                        key: fromUtf8(attribute.key),
                        value: fromUtf8(attribute.value),
                    })
                ),
            })
        )
    }

    public async getEndBlockEvents(height: number): Promise<StringEvent[]> {
        const results: BlockResultsResponse =
            await this.myTmClient.blockResults(height)
        return this.convertTendermintEvents(results.endBlockEvents)
    }
}
