import c, { UploadApiErrorResponse, UploadApiResponse } from "cloudinary"

const cloudinary = c.v2;

export const uploads = (
    file: string,
    publicId?: string,
    overwrite?: boolean,
    invalidate?: boolean,
): Promise<UploadApiErrorResponse | UploadApiResponse> => {
    return new Promise((resolve) => {
        cloudinary.uploader.upload(file, { publicId, overwrite, invalidate, resource_type: 'auto' },
            (error: UploadApiErrorResponse | undefined, result: UploadApiResponse | undefined) => {
                if (error) resolve(error);
                if (result) resolve(result);
            })
    })
}

export const videoUpload = (
    file: string,
    publicId?: string,
    overwrite?: boolean,
    invalidate?: boolean,
): Promise<UploadApiErrorResponse | UploadApiResponse> => {
    return new Promise((resolve) => {
        cloudinary.uploader.upload(file, { publicId, overwrite, invalidate, chunk_size: 50000, resource_type: 'video' },
            (error: UploadApiErrorResponse | undefined, result: UploadApiResponse | undefined) => {
                if (error) resolve(error);
                if (result) resolve(result);
            })
    })
} 