import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatCardModule } from "@angular/material/card";
import { MatButtonModule } from "@angular/material/button";
import { MatProgressBarModule } from "@angular/material/progress-bar";
import { MatChipsModule } from "@angular/material/chips";
import { MatExpansionModule } from "@angular/material/expansion";
import { MatDividerModule } from "@angular/material/divider";
import { MatIconModule } from "@angular/material/icon";
import { MatTooltipModule } from "@angular/material/tooltip";

const modules = [
  MatCardModule,
  MatButtonModule,
  MatProgressBarModule,
  MatChipsModule,
  MatExpansionModule,
  MatDividerModule,
  MatIconModule,
  MatTooltipModule,
]

@NgModule({
  imports: [CommonModule, ...modules],
  exports: [...modules]
})
export class MaterialModule {
}
