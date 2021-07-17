import { Component, OnInit } from '@angular/core';
import { UntilDestroy, untilDestroyed } from '@ngneat/until-destroy';
import { finalize } from 'rxjs/operators';
import { Stream } from './shared/models/stream.model';
import { APIService } from './shared/providers/api.service';

@UntilDestroy()
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html'
})
export class AppComponent implements OnInit {
  streams: Stream[] = [];
  loading = false;

  constructor(private api: APIService) {}

  ngOnInit(): void {
    this.findStreams();
  }

  findStreams(): void {
    this.loading = true;
    this.api
      .findStreams()
      .pipe(
        finalize(() => (this.loading = false)),
        untilDestroyed(this)
      )
      .subscribe(resp => {
        this.streams = resp;
      });
  }
}
