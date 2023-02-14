import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { Movie, MovieDocument } from './entities/movies';

@Injectable()
export class AppService {
  constructor(@InjectModel(Movie.name) private model: Model<MovieDocument>) {}

  async findByIds(ids: string[]): Promise<Movie[]> {
    return this.model.find({ _id: { $in: ids } });
  }
}
