import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { APIService } from './providers/api.service';
import { ModalModule } from './components/modal/modal.module';

@NgModule({
  imports: [CommonModule, HttpClientModule, ModalModule],
  declarations: [],
  providers: [APIService],
  exports: [ModalModule]
})
export class SharedModule {}
