import { createHotelDTO } from "../dto/hotel.dto";
import { HotelRepository } from "../repositories/hotel.repository";
import { BadRequestError } from "../utils/errors/app.error";


const hotelRepository=new HotelRepository();

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
    const hotel=await hotelRepository.create(hotelData);
    return hotel;
}

export async function getHotelByService(id:number) {
    const hotel=await hotelRepository.findById(id);
    return hotel;
}

export async function getAllHotelsService() {
    const hotel=await hotelRepository.findAll();
    return hotel;
}

export async function deleteHotelService(id:number) {
    const response=await hotelRepository.softDelete(id);
    return response;
}