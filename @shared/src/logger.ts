import w, { Logger, transport } from 'winston';
import { ElasticsearchTransformer, ElasticsearchTransport, LogData, TransformedData } from 'winston-elasticsearch'

const esTransformer = (logData: LogData): TransformedData => {
    return ElasticsearchTransformer(logData)
}
export const winstonLogger = (elasticsearchNode: string, name: string, level: string) => {
    const options = {
        console: {
            level,
            handleExceptions: true,
            json: false,
            colorize: true,
        },
        elasticsearch: {
            level,
            transformer: esTransformer,
            clientOptions: {
                node: elasticsearchNode,
                log: level,
                maxRetries: 2,
                requestTimeout: 10000,
                sniffOnStart: false,
            }
        }
    }

    const esTransport: ElasticsearchTransport = new ElasticsearchTransport(options.elasticsearch);
    
    const logger: Logger = w.createLogger({
        exitOnError: false,
        defaultMeta: { service: name },
        transports: [new w.transports.Console(options.console), esTransport]
    })
}