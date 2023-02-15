import { Expose } from 'class-transformer';

export class MovieResponseDto {
  @Expose()
  title: string;

  @Expose()
  year: number;
}
