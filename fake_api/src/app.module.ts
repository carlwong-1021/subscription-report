import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { Customer } from './entities/customer';

@Module({
  imports: [
    TypeOrmModule.forRoot({
      type: 'mysql',
      host: 'localhost',
      port: 3306,
      username: 'demo',
      password: 'demo',
      database: 'classicmodels',
      entities: [Customer],
      logging: true,
    }),
    TypeOrmModule.forFeature([Customer]),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
