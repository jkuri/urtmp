import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { APIService } from './providers/api.service';
import { ModalModule } from './components/modal/modal.module';
import { SocketService } from './providers/socket.service';
import { DataService } from './providers/data.service';

@NgModule({
  imports: [CommonModule, HttpClientModule, ModalModule],
  providers: [APIService, SocketService, DataService],
  exports: [ModalModule]
})
export class SharedModule {}
