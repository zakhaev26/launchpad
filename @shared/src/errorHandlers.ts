import { StatusCodes } from "http-status-codes"

export interface IErrorResponse {
    message: string;
    statusCode: number;
    status: string;
    comingFrom: string;
    serializeErrors(): IError;
}


export interface IError {
    message: string;
    statusCode: number;
    status: string;
    comingFrom: string;
}

export abstract class CustomError extends Error {
    abstract statusCode: number;
    abstract status: string;
    comingFrom: string;

    constructor(message: string, comingFrom: string) {
        super(message);
        this.comingFrom = comingFrom;
    }

    serializeErrors(): IError {
        return {
            message: this.message,
            comingFrom: this.comingFrom,
            status: this.status,
            statusCode: this.statusCode,
        }
    }
}

export class BadRequestError extends CustomError {
    statusCode = StatusCodes.BAD_REQUEST;
    status = 'error';

    constructor(message: string, comingFrom: string) {
        super(message, comingFrom);
        }
}