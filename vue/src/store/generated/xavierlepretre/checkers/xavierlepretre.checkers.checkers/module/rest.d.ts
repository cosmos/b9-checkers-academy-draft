export interface CheckersMsgCreateGameResponse {
    idValue?: string;
}
export interface CheckersMsgPlayMoveResponse {
    idValue?: string;
    /** @format int64 */
    capturedX?: string;
    /** @format int64 */
    capturedY?: string;
    winner?: string;
}
export declare type CheckersMsgRejectGameResponse = object;
export interface CheckersNextGame {
    creator?: string;
    /** @format uint64 */
    idValue?: string;
    fifoHead?: string;
    fifoTail?: string;
}
export interface CheckersQueryAllStoredGameResponse {
    StoredGame?: CheckersStoredGame[];
    /**
     * PageResponse is to be embedded in gRPC response messages where the
     * corresponding request message has used PageRequest.
     *
     *  message SomeResponse {
     *          repeated Bar results = 1;
     *          PageResponse page = 2;
     *  }
     */
    pagination?: V1Beta1PageResponse;
}
export interface CheckersQueryGetNextGameResponse {
    NextGame?: CheckersNextGame;
}
export interface CheckersQueryGetStoredGameResponse {
    StoredGame?: CheckersStoredGame;
}
export interface CheckersStoredGame {
    creator?: string;
    index?: string;
    game?: string;
    turn?: string;
    red?: string;
    black?: string;
    moveCount?: string;
    beforeId?: string;
    afterId?: string;
    deadline?: string;
}
export interface ProtobufAny {
    "@type"?: string;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
    /**
     * key is a value returned in PageResponse.next_key to begin
     * querying the next page most efficiently. Only one of offset or key
     * should be set.
     * @format byte
     */
    key?: string;
    /**
     * offset is a numeric offset that can be used when key is unavailable.
     * It is less efficient than using key. Only one of offset or key should
     * be set.
     * @format uint64
     */
    offset?: string;
    /**
     * limit is the total number of results to be returned in the result page.
     * If left empty it will default to a value to be set by each app.
     * @format uint64
     */
    limit?: string;
    /**
     * count_total is set to true  to indicate that the result set should include
     * a count of the total number of items available for pagination in UIs.
     * count_total is only respected when offset is used. It is ignored when key
     * is set.
     */
    countTotal?: boolean;
}
/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
    /** @format byte */
    nextKey?: string;
    /** @format uint64 */
    total?: string;
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title checkers/genesis.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryNextGame
     * @summary Queries a nextGame by index.
     * @request GET:/xavierlepretre/checkers/checkers/nextGame
     */
    queryNextGame: (params?: RequestParams) => Promise<HttpResponse<CheckersQueryGetNextGameResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryStoredGameAll
     * @summary Queries a list of storedGame items.
     * @request GET:/xavierlepretre/checkers/checkers/storedGame
     */
    queryStoredGameAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<CheckersQueryAllStoredGameResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryStoredGame
     * @summary Queries a storedGame by index.
     * @request GET:/xavierlepretre/checkers/checkers/storedGame/{index}
     */
    queryStoredGame: (index: string, params?: RequestParams) => Promise<HttpResponse<CheckersQueryGetStoredGameResponse, RpcStatus>>;
}
export {};
