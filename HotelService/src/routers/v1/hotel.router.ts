import express from "express"
import { createHotelHandler, deleteHotelHandler, getAllHotelsHandler, getHotelByHandler } from "../../controllers/hotel.controller"
import { hotelSchema } from "../../validators/hotel.validator";
import { validateRequestBody } from "../../validators";

const hotelRouter=express.Router()

hotelRouter.post('/',validateRequestBody(hotelSchema),createHotelHandler); // TODO: Release this TS compilation issue

hotelRouter.get('/:id',getHotelByHandler); // TODO: Release this TS compilation issue

hotelRouter.get('/',getAllHotelsHandler)

hotelRouter.delete('/:id',deleteHotelHandler)

export default hotelRouter;