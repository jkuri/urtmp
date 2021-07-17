import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { APIService } from './providers/api.service';

@NgModule({
  imports: [CommonModule, HttpClientModule],
  declarations: [],
  providers: [APIService],
  exports: []
})
export class SharedModule {}
