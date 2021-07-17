import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppComponent } from './app.component';
import { VideoModalComponent } from './components/video-modal/video-modal.component';
import { SharedModule } from './shared/shared.module';
import { StreamListItemComponent } from './components/stream-list-item/stream-list-item.component';

@NgModule({
  declarations: [AppComponent, VideoModalComponent, StreamListItemComponent],
  imports: [BrowserModule, SharedModule],
  bootstrap: [AppComponent]
})
export class AppModule {}
