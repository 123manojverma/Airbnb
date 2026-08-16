import { Prisma } from "../generated/prisma/client";
import prisma from "../prisma/client";

export async function createBooking(bookingInput:Prisma.BookingCreateInput) {
    const booking=await prisma.booking.create({
        data:bookingInput
    })

    return booking;
}