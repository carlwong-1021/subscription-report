import { Controller, Get, ParseArrayPipe, Query } from '@nestjs/common';
import { plainToClass } from 'class-transformer';
import { AppService } from './app.service';
import { MovieResponseDto } from './movie/movie.response.dto';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get('/')
  async getCustomer(
    @Query('ids', new ParseArrayPipe({ items: String, separator: ',' }))
    ids: string[],
  ) {
    return plainToClass(MovieResponseDto, this.appService.findByIds(ids), {
      strategy: 'excludeAll',
    });
  }
}
