import { Entity, Column, PrimaryGeneratedColumn } from 'typeorm';

@Entity({
  name: 'customers',
})
export class Customer {
  @PrimaryGeneratedColumn()
  customerNumber: number;

  @Column()
  customerName: string;

  @Column()
  phone: string;

  @Column()
  city: string;
}
