import { IdempotencyKey, Prisma } from "../generated/prisma/client";
import {prisma} from "../prisma/client";
import { validate as isValidUUID } from "uuid";
import { InternalServerError, NotFoundError } from "../utils/errors/app.error";

export async function createBooking(bookingInput: Prisma.BookingCreateInput) {
    const booking = await prisma.booking.create({
        data: bookingInput
    })

    return booking;
}

export async function createIdempotencyKey(key: string, bookingId: number) {
    const idempotencyKey = await prisma.idempotencyKey.create({
        data: {
            idemKey:key,
            booking: {
                connect: { id: bookingId }
            }
        }
    })

    return idempotencyKey;
}

export async function getIdempotencyKeyWithLock(tx: Prisma.TransactionClient, key: string) {

    if(!isValidUUID(key)){
        throw new InternalServerError("Invalid idempotency key format")
    }

    const idempotencyKey:Array<IdempotencyKey> = await tx.$queryRaw(
        Prisma.raw(`SELECT * FROM idempotencyKey WHERE idemKey='${key}' FOR UPDATE;`)
    )

    console.log("Idempotency key with lock:",idempotencyKey);

    if(!idempotencyKey || idempotencyKey.length==0){
        throw new NotFoundError("Idempotency key not found")
    }

    return idempotencyKey[0];
}

export async function getBookingById(bookingId: number) {
    const booking = await prisma.booking.findUnique({
        where: {
            id: bookingId
        }
    })
    return booking;
}

export async function confirmBooking(tx:Prisma.TransactionClient,bookingId:number) {
    const booking=await tx.booking.update({
        where:{
            id:bookingId
        },
        data:{
            status:"CONFIRMED"
        }
    })

    return booking;
}

export async function cancelBooking(bookingId:number) {
    const booking=await prisma.booking.update({
        where:{
            id:bookingId
        },
        data:{
            status:"CANCELLED"
        }
    })

    return booking;
}

export async function finalizeIdempotencyKey(tx:Prisma.TransactionClient,key:string) {
    const idempotencyKey=await tx.idempotencyKey.update({
        where:{
            idemKey:key
        },
        data:{
            finalized:true
        }
    })

    return idempotencyKey;
}

