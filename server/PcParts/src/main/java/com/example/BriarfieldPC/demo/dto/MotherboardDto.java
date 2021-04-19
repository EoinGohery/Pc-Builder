package com.example.BriarfieldPC.demo.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class MotherboardDto {
    public long Id;
    public String Manufacturer;
    public String Name;
    public Integer RamSlots;
    public Integer DriveSlots;
    public Integer MaxRam;
    public String Socket;
    public Integer Tdp;
    public Integer Price;
}
