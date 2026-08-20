import { Job, Worker } from "bullmq";
import { NotificationDto } from "../dto/notification.dto";
import { MAILER_QUEUE } from "../queues/maller.queue";
import { getRedisConnObject } from "../config/redis.config";
import { MAILER_PAYLOAD } from "../producers/email.producer";
import { renderMailTemplate } from "../templates/templates.handler";
import { sendEmail } from "../services/mailer.service";
import logger from "../config/logger.config";

export const setupMailerWorker=()=>{
    const emailProcessor=new Worker<NotificationDto>(
        MAILER_QUEUE, // Name of the queue
        async(job:Job)=>{
            if(job.name!==MAILER_PAYLOAD){
                throw new Error("Invalid job name")
            }
    
            // call the service layer from here
            const payload=job.data;
            console.log(`Processing eamil for: ${JSON.stringify(payload)}`);

            const emailContent=await renderMailTemplate(payload.templateId,payload.params);

            await sendEmail(payload.to,payload.subject,emailContent)

            logger.info(`Email sent to ${payload.id} with subject ${payload.subject}`)
        }, //Process function
        { 
            connection:getRedisConnObject()
        }
    )
    
    emailProcessor.on("failed",()=>{
        console.log("Email processing failedd");
    })
    
    emailProcessor.on("completed",()=>{
        console.log("Email processing completed successfully");
    })
}
