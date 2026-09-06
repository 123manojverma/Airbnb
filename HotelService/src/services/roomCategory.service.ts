import { createRoomCategoryDTO } from "../dto/roomCategory.dto";
import { HotelRepository } from "../repositories/hotel.repository";
import RoomCategoryRepository from "../repositories/roomCategory.repository";
import { NotFoundError } from "../utils/errors/app.error";

const roomCategoryRepository=new RoomCategoryRepository();
const hotelRepository=new HotelRepository();

export async function createRoomCategoryService(roomCategorydata:createRoomCategoryDTO) {
    const roomCategory=await roomCategoryRepository.create(roomCategorydata);
    return roomCategory;
}

export async function getRoomCategoryByIdService(id:number) {
    const roomCategory=await roomCategoryRepository.findById(id);
    return roomCategory;
}

export async function getAllRoomCategoriesByHotelService(hotelId:number) {
    // Check if the hotel exist or not
    const hotel=await hotelRepository.findById(hotelId);

    if(!hotel){
        throw new NotFoundError(`Hotel with id ${hotelId} not found`);
    }

    // find all room categories by hotel id
    const roomCategories=await roomCategoryRepository.findAllByHotelId(hotelId);

    return roomCategories;
}

export async function deleteRoomCategoryService(id:number) {
    const roomCategory=await roomCategoryRepository.findById(id);

    if(!roomCategory){
        throw new NotFoundError(`Room category with id ${id} not found`);
    }

    await roomCategoryRepository.delete({id});
    return true;
}