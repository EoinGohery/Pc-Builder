package com.example.BriarfieldPC.demo.service;

import com.example.BriarfieldPC.demo.dto.*;
import com.example.BriarfieldPC.demo.entity.*;
import com.example.BriarfieldPC.demo.repository.*;
import lombok.AllArgsConstructor;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@Service
@AllArgsConstructor
public class PcService {
    @Autowired
    private CpuRepositroy cpuRepository;
    @Autowired
    private GpuRepository gpuRepository;
    @Autowired
    private DriveRepository driveRepository;
    @Autowired
    private MotherboardRepository motherboardRepository;
    @Autowired
    private PsuRepository psuRepository;
    @Autowired
    private RamRepository ramRepository;

    public List<CpuDto> getAllCpu(){
        List<Cpu> cpus;
        List<CpuDto> dtos =  new ArrayList<CpuDto>();
         cpus = cpuRepository.findAll();
         for(int i = 0; i < cpus.size(); i++){
             dtos.add(new CpuDto(cpus.get(i).getId(),cpus.get(i).getManufacturer(),cpus.get(i).getName(),cpus.get(i).getCores(),cpus.get(i).getClock(),cpus.get(i).getTdp(),cpus.get(i).getSocket(),cpus.get(i).getPrice()));
         }
         return dtos;
    }
    public List<GpuDto> getAllGpu(){
        List<Gpu> gpus;
        List<GpuDto> dtos =  new ArrayList<GpuDto>();
        gpus = gpuRepository.findAll();
        for(int i = 0; i < gpus.size(); i++){
            dtos.add(new GpuDto(gpus.get(i).Id,gpus.get(i).Manufacturer,gpus.get(i).GpuName,gpus.get(i).Memory,gpus.get(i).Clock,gpus.get(i).Tdp,gpus.get(i).Price));
        }
        return dtos;
    }
    public List<DriveDto> getAllDrivers(){

        List<Drive> drives;
        List<DriveDto> dtos =  new ArrayList<DriveDto>();
        drives = driveRepository.findAll();
        for(int i = 0; i < drives.size(); i++){
            dtos.add(new DriveDto(drives.get(i).Id,drives.get(i).Manufacturer,drives.get(i).Name,drives.get(i).Size,drives.get(i).Tdp,drives.get(i).Price,drives.get(i).technology));
        }
        return dtos;
    }
    public List<MotherboardDto> getAllMotherboards(){

        List<Motherboard> motherboards;
        List<MotherboardDto> dtos =  new ArrayList<MotherboardDto>();
        motherboards = motherboardRepository.findAll();
        for(int i = 0; i < motherboards.size(); i++){
            dtos.add(new MotherboardDto(motherboards.get(i).Id,motherboards.get(i).Manufacturer,motherboards.get(i).Name,motherboards.get(i).RamSlots,motherboards.get(i).DriveSlots,motherboards.get(i).MaxRam,motherboards.get(i).Socket,motherboards.get(i).Tdp,motherboards.get(i).Price));
        }
        return dtos;
    }
    public List<PSUDto> getAllPSU(){
        List<PSU> psus;
        List<PSUDto> dtos =  new ArrayList<PSUDto>();
        psus = psuRepository.findAll();
        for(int i = 0; i < psus.size(); i++){
            dtos.add(new PSUDto(psus.get(i).Id,psus.get(i).Manufacturer,psus.get(i).Name,psus.get(i).Price,psus.get(i).Capacity,psus.get(i).Rating));
        }
        return dtos;
    }
    public List<RamDto> getAllRam(){
        List<Ram> ram;
        List<RamDto> dtos =  new ArrayList<RamDto>();
        ram = ramRepository.findAll();
        for(int i = 0; i < ram.size(); i++){
            dtos.add(new RamDto(ram.get(i).Id,ram.get(i).Manufacturer,ram.get(i).Name,ram.get(i).Memory,ram.get(i).Clock,ram.get(i).Tdp,ram.get(i).Price));
        }
        return dtos;
    }
}
