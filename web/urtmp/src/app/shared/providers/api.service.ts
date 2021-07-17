import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { Stream } from '../models/stream.model';

@Injectable({ providedIn: 'root' })
export class APIService {
  constructor(private http: HttpClient) {}

  findStreams(): Observable<Stream[]> {
    return this.http
      .get<Stream[]>('/api/v1/streams')
      .pipe(
        map(resp => (resp && resp.length ? resp.map(r => new Stream(r.key, r.subscribers)) : []))
      );
  }
}
