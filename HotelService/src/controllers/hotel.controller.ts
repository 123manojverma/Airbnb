import { Request, Response, NextFunction } from "express";
import { createHotelService, getHotelByService, getAllHotelsService, deleteHotelService } from "../services/hotel.service";
import { BadRequestError } from "../utils/errors/app.error";
import {StatusCodes} from "http-status-codes"

export async function createHotelHandler(req: Request, res: Response, next: NextFunction) {
    // 1. Call the service layer

    const hotelResponse = await createHotelService(req.body);

    // 2. Send the response

    res.status(StatusCodes.CREATED).json({
        message: "Hotel created successfully",
        data: hotelResponse,
        success: true
    })

    // try{
    //     const hotelData: createHotelDTO=req.body;
    //     const hotel=await createHotelService(hotelData);
    //     res.status(201).json(hotel)
    // }catch(err){
    //     next(err)
    // }
}

export async function getHotelByHandler(req: Request, res: Response, next: NextFunction) {
    try {
        const id = Number(req.params.id);
        if (isNaN(id) || id <= 0) {
            throw new BadRequestError("Invalid hotel ID");
        }
        const hotelResponse = await getHotelByService(id);

        res.status(StatusCodes.OK).json({
            message: "Hotel found successfully",
            data: hotelResponse,
            success: true
        })
    } catch(err) {
        next(err);
    }
}

export async function getAllHotelsHandler(req: Request, res: Response, next: NextFunction) {
    try {
        const hotels = await getAllHotelsService(); 
        res.status(StatusCodes.OK).json({
            message: "Hotels retrieved successfully",
            hotels,
            success:true
        })
    } catch(err) {
        next(err);  // Pass errors to error middleware
    }
}

export async function deleteHotelHandler(req: Request, res: Response, next: NextFunction) {
    try {
        const hotelsresponse = await deleteHotelService(Number(req.params.id)); 
        res.status(StatusCodes.OK).json({
            message: "Hotels retrieved successfully",
            data:hotelsresponse,
            success:true
        })
    } catch(err) {
        next(err); 
    }
}

export async function updateHotelHandler(req: Request, res: Response, next: NextFunction) {
    res.status(StatusCodes.NOT_IMPLEMENTED)
}