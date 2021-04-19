package com.example.BriarfieldPC.demo.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class DriveDto {
    public long Id;
    public String Manufacturer;
    public String Name;
    public Integer Size;
    public Integer Tdp;
    public Integer Price;
    public String technology;
}
