package com.example.BriarfieldPC.demo.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class RamDto {
    public long Id;
    public String Manufacturer;
    public String Name;
    public Integer Memory;
    public String Clock;
    public Integer Tdp;
    public Integer Price;
}
