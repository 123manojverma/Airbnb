import { CreationOptional, InferAttributes, InferCreationAttributes, Model } from "sequelize";
import sequelize from "./sequelize";

class Hotel extends Model<InferAttributes<Hotel>, InferCreationAttributes<Hotel>> {
    declare id: CreationOptional<number>;
    declare name: string;
    declare address: string;
    declare location: string;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
    declare deleted_at: CreationOptional<Date | null>;
    declare rating?: number;
    declare rating_count?: number;
}

Hotel.init({
    id: {
        type: 'INTEGER',
        autoIncrement: true,
        primaryKey: true,
    },
    name: {
        type: 'VARCHAR(255)',
        allowNull: false
    },
    address: {
        type: 'STRING',
        allowNull: false
    },
    location: {
        type: 'VARCHAR(255)',
        allowNull: false
    },
    createdAt: {
        type: 'TIMESTAMP',
        defaultValue: new Date(),
    },
    updatedAt: {
        type: 'TIMESTAMP',
        defaultValue: new Date(),
    },
    deleted_at:{
        type:'DATE',
        defaultValue:null
    },
    rating: {
        type: 'FLOAT',
        defaultValue: null
    },
    rating_count: {
        type: 'INTEGER',
        defaultValue: null
    }
}, {
    tableName:"hotels",
    sequelize: sequelize,
    // underscored:true, // use snake_case for column names
    timestamps:true //createdAt,updatedAt
})

export default Hotel;

// Hotel.create({name:"",address:"",location:""})