import { Component, Input, OnInit } from '@angular/core';
import { ModalService } from 'src/app/shared/components/modal/modal.service';
import { Stream } from 'src/app/shared/models/stream.model';
import { VideoModalComponent } from '../video-modal/video-modal.component';

@Component({
  selector: 'app-stream-list-item',
  templateUrl: './stream-list-item.component.html'
})
export class StreamListItemComponent implements OnInit {
  @Input() stream!: Stream;

  constructor(private modal: ModalService) {}

  ngOnInit(): void {}

  openVideoModal(): void {
    const modal = this.modal.open(VideoModalComponent, { size: 'large' });
    modal.componentInstance.key = this.stream.key;
  }
}
