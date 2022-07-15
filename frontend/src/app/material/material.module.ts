import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatCardModule } from "@angular/material/card";
import { MatButtonModule } from "@angular/material/button";
import { MatProgressBarModule } from "@angular/material/progress-bar";

const modules = [
  MatCardModule,
  MatButtonModule,
  MatProgressBarModule,
]

@NgModule({
  imports: [CommonModule, ...modules],
  exports: [...modules]
})
export class MaterialModule {
}
