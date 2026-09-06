import { RoomType } from "../db/models/roomCategory";

export type createRoomCategoryDTO={
    hotelId:number;
    price:number;
    roomType:RoomType;
    roomCount:number
}