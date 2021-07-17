import { Component, OnDestroy, OnInit } from '@angular/core';
import { UntilDestroy, untilDestroyed } from '@ngneat/until-destroy';
import { filter, finalize } from 'rxjs/operators';
import { Stream } from './shared/models/stream.model';
import { APIService } from './shared/providers/api.service';
import { DataService } from './shared/providers/data.service';

@UntilDestroy()
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html'
})
export class AppComponent implements OnInit, OnDestroy {
  streams: Stream[] = [];
  loading = false;

  constructor(private api: APIService, private data: DataService) {}

  ngOnInit(): void {
    this.findStreams();

    this.data.socketOutput
      .pipe(
        filter(e => e.type === '/events'),
        untilDestroyed(this)
      )
      .subscribe(() => {
        this.findStreams();
      });

    this.data.subscribeToEvent(`/events`);
  }

  ngOnDestroy(): void {
    this.data.unsubscribeAll();
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
