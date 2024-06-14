import JWT from 'jsonwebtoken';
import { Request, Response, NextFunction } from "express";
import { NotAuthorizedError } from './errorHandlers';

const tokens: Array<string> = ['auth', 'seller', 'gig', 'search', 'buyer', 'message', 'order', 'review'];

export function verifyGatewayRequest(req: Request, res: Response, nex: NextFunction): void {

    if (!req.headers?.gatewayToken)
        throw new NotAuthorizedError('Invalid Request', 'verifyGatewayRequest() method: Request not coming from api gateway');

    const token: string = req.headers.gatewayToken as string;

    if (!token)
        throw new NotAuthorizedError('Invalid Request', 'verifyGatewayRequest() method: Request not coming from api gateway');

    try {
        const payload: {
            id: string,
            iat: number,
        } = JWT.verify(token, '') as { id: string, iat: number, }

        if (!tokens.includes(payload.id))
            throw new NotAuthorizedError('Invalid Request', 'verifyGatewayRequest() method: Request not coming from api gateway');

        
    } catch (error: any) {
        throw new NotAuthorizedError('Invalid Request', 'verifyGatewayRequest() method: Request not coming from api gateway');
    }

}