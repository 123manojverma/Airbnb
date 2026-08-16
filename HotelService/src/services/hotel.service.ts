import { createHotelDTO } from "../dto/hotel.dto";
import { createHotel, getAllHotels, getHotelById, softDeleteHotel } from "../repositories/hotel.repository";
import { BadRequestError } from "../utils/errors/app.error";

const blockListedAddresses=[
    "123 Fake St",
    "456 Elm St",
    "789 Maple Ave",
]

export function isAddressBlockListed(address:string):boolean{
    return blockListedAddresses.includes(address);
}

export async function createHotelService(hotelData:createHotelDTO) {
    if(isAddressBlockListed(hotelData.address)){
        throw new BadRequestError("Address is blocklisted");
    }
    const hotel=await createHotel(hotelData);
    return hotel;
}

export async function getHotelByService(id:number) {
    const hotel=await getHotelById(id);
    return hotel;
}

export async function getAllHotelsService() {
    const hotel=await getAllHotels();
    return hotel;
}

export async function deleteHotelService(id:number) {
    const response=await softDeleteHotel(id);
    return response;
}