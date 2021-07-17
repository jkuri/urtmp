import { Injectable, Injector } from '@angular/core';

export type ModalOption =
  | 'backdrop'
  | 'backdropOpacity'
  | 'beforeDismiss'
  | 'container'
  | 'injector'
  | 'keyboard'
  | 'scrollable'
  | 'size';

export interface ModalOptions {
  backdrop?: boolean | 'white';
  backdropOpacity?: number;
  beforeDismiss?: () => boolean | Promise<boolean>;
  container?: string;
  injector?: Injector;
  keyboard?: boolean;
  scrollable?: boolean;
  size?: 'small' | 'large' | 'medium' | 'fullscreen';
}

@Injectable({ providedIn: 'root' })
export class ModalConfig implements Required<ModalOptions> {
  backdrop: boolean | 'white' = true;
  backdropOpacity = 0.8;
  beforeDismiss!: () => boolean | Promise<boolean>;
  container!: string;
  injector!: Injector;
  keyboard = true;
  scrollable!: boolean;
  size!: 'small' | 'large' | 'medium';
}
