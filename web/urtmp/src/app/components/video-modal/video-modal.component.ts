import { Component, ElementRef, Input, OnDestroy, OnInit } from '@angular/core';
import { ActiveModal } from 'src/app/shared/components/modal/modal-ref.class';
import { default as mpegts } from 'mpegts.js';
import { UntilDestroy, untilDestroyed } from '@ngneat/until-destroy';
import { timer } from 'rxjs';

@UntilDestroy()
@Component({
  selector: 'app-video-modal',
  templateUrl: './video-modal.component.html',
  styleUrls: ['./video-modal.component.sass']
})
export class VideoModalComponent implements OnInit, OnDestroy {
  @Input() key: string = '';

  el!: HTMLMediaElement;
  player!: mpegts.Player;

  constructor(private elementRef: ElementRef, public activeModal: ActiveModal) {}

  ngOnInit(): void {
    this.el = this.elementRef.nativeElement.querySelector('.player');
    const source = `/api/v1/live/${this.key}`;

    this.player = mpegts.createPlayer(
      {
        type: 'flv',
        url: source,
        isLive: true,
        hasAudio: false,
        hasVideo: true,
        cors: true
      },
      {
        enableStashBuffer: false,
        isLive: true,
        lazyLoad: false,
        deferLoadAfterSourceOpen: false,
        stashInitialSize: 128,
        liveBufferLatencyChasing: true
      }
    );

    this.player.attachMediaElement(this.el);
    this.player.load();
    this.player.play();

    timer(100, 1000)
      .pipe(untilDestroyed(this))
      .subscribe(() => {
        const buffered = this.el.buffered;
        if (buffered.length > 0) {
          const end = buffered.end(0);
          if (end - this.el.currentTime > 0.15) {
            this.el.currentTime = end - 0.1;
          }
        }
      });
  }

  ngOnDestroy(): void {
    this.player.destroy();
  }
}
