import { CreateBookingDTO } from "../dto/booking.dto";
import { confirmBooking, createBooking, createIdempotencyKey, finalizeIdempotencyKey, getIdempotencyKeyWithLock } from "../repositories/booking.repository";
import { BadRequestError, InternalServerError, NotFoundError } from "../utils/errors/app.error";
import { generateIdempotencyKey } from "../utils/generateIdempotencyKey";
import {prisma} from '../prisma/client'
import { redlock } from "../config/redis.config";
import { serverConfig } from "../config";

export async function createBookingService(
    createBookingDTO:CreateBookingDTO
) {
    const ttl=serverConfig.LOCK_TTL
    const bookingResource=`hotel:${createBookingDTO.hotelId}`

    let lock;

    try {
        lock = await redlock.acquire([bookingResource], ttl);
        console.log(lock);
        const booking = await createBooking({
            userId:createBookingDTO.userId,
            hotelId:createBookingDTO.hotelId,
            totalGuests:createBookingDTO.totalGuests,
            bookingAmount:createBookingDTO.bookingAmount
        });

        const idempotencyKey = generateIdempotencyKey();

        await createIdempotencyKey(idempotencyKey, booking.id);

        return {
            bookingId: booking.id,
            idempotencyKey
        };
    }catch(err){
        throw new InternalServerError("Failed to acquire lock for booking resource")
    }

}

export async function confirmBookingService(idempotencyKey:string) {

    return await prisma.$transaction(async(tx)=>{
        const idempotencyKeyData=await getIdempotencyKeyWithLock(tx,idempotencyKey)
    
        if(!idempotencyKeyData){
            throw new NotFoundError('Idempotency key not found');
        }
    
        if(idempotencyKeyData.finalized){
            throw new BadRequestError('Idempotency key already finalized');
        }
    
        const booking=await confirmBooking(tx,idempotencyKeyData.bookingId);
        await finalizeIdempotencyKey(tx,idempotencyKey)
    
        return booking;
    })

}