import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { In, Repository } from 'typeorm';
import { Customer } from './entities/customer';

@Injectable()
export class AppService {
  constructor(
    @InjectRepository(Customer) private customerRepo: Repository<Customer>,
  ) {}

  async findByCustomerIds(ids: number[]): Promise<Customer[]> {
    const customers = await this.customerRepo.find({
      where: {
        customerNumber: In(ids),
      },
    });
    console.log(customers);
    return customers;
  }
}
